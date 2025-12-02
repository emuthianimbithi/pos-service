package services

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
)

// StorageProvider defines the interface for file storage
type StorageProvider interface {
	UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error)
	DeleteFile(path string) error
}

// LocalStorage implements StorageProvider for local filesystem
type LocalStorage struct {
	BaseURL   string
	UploadDir string
}

func NewLocalStorage(baseURL, uploadDir string) *LocalStorage {
	// Ensure upload directory exists
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.MkdirAll(uploadDir, 0755)
	}
	return &LocalStorage{
		BaseURL:   baseURL,
		UploadDir: uploadDir,
	}
}

func (s *LocalStorage) UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	// Create folder if it doesn't exist
	targetDir := filepath.Join(s.UploadDir, folder)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		os.MkdirAll(targetDir, 0755)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	path := filepath.Join(targetDir, filename)

	// Create destination file
	dst, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy content
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	// Return public URL
	return fmt.Sprintf("%s/uploads/%s/%s", s.BaseURL, folder, filename), nil
}

func (s *LocalStorage) DeleteFile(path string) error {
	// Extract relative path from URL or use as is if it's a file path
	// This is a simplified implementation
	return os.Remove(path)
}

// GCPStorage implements StorageProvider for Google Cloud Storage
type GCPStorage struct {
	BucketName string
	Client     *storage.Client
}

func NewGCPStorage(bucketName string) (*GCPStorage, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &GCPStorage{
		BucketName: bucketName,
		Client:     client,
	}, nil
}

func (s *GCPStorage) UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	ctx := context.Background()

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s/%s%s", folder, uuid.New().String(), ext)

	wc := s.Client.Bucket(s.BucketName).Object(filename).NewWriter(ctx)
	wc.ContentType = header.Header.Get("Content-Type")

	if _, err := io.Copy(wc, file); err != nil {
		return "", err
	}
	if err := wc.Close(); err != nil {
		return "", err
	}

	// Return public URL
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", s.BucketName, filename), nil
}

func (s *GCPStorage) DeleteFile(path string) error {
	// TODO: Parse object name from URL
	return nil
}

// StorageService is the main service that uses a provider
type StorageService struct {
	provider StorageProvider
}

func NewStorageService(providerType string) (*StorageService, error) {
	var provider StorageProvider
	var err error

	switch providerType {
	case "gcp":
		bucket := os.Getenv("GCP_BUCKET_NAME")
		if bucket == "" {
			return nil, fmt.Errorf("GCP_BUCKET_NAME environment variable is required")
		}
		provider, err = NewGCPStorage(bucket)
	case "local":
		baseURL := os.Getenv("APP_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}
		provider = NewLocalStorage(baseURL, "uploads")
	default:
		// Default to local
		baseURL := os.Getenv("APP_URL")
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}
		provider = NewLocalStorage(baseURL, "uploads")
	}

	if err != nil {
		return nil, err
	}

	return &StorageService{provider: provider}, nil
}

func (s *StorageService) Upload(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	return s.provider.UploadFile(file, header, folder)
}
