package provider

import "fmt"

// Provider 数据提供接口
type provider interface {

	// AuthWithUsernameAndPassword 使用用户名密码认证
	AuthWithUsernameAndPassword(username string, password string)

	// GetAllCategories 获取所有的分类信息
	GetAllCategories()
}

const (
	PROVIDER_WORDPRESS = iota
)

// NewProvider 创建一个Provider
func NewProvider(providerType int) *provider {
	switch providerType {
	case PROVIDER_WORDPRESS:
		fmt.Println("---")
		return nil
	default:
		return nil
	}
}
