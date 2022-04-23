# 路由

## 版本及目录

一般每个大版本建立一个单独的目录，以便于管理，目录名同版本号，如 `v1` 版本的 API 的声明，都在 `v1` 目录下的文件中。

## 声明路由

### 初始化

约定：每个版本下应有一个 `init.go` 文件，其中定义以下变量：

- `apiVersion` 版本号，如 `v1`
- `ApiRouter` 路由实例，构建于 `components.router`
- `LoadRoutes []func(group *gin.RouterGroup)` 切片，元素为用于执行载入路由的函数

### 定义路由

为了方便分组管理，可以按功能模块、PATH 等维度，将路由分为多个块并存放于不同的文件中。

所有路由都应在函数 `func (group *gin.RouterGroup)` 中定义，且该函数可公开访问，并将该函数名放到 `LoadRoutes` 变量中，才会自动载入。

所有路由都使用 `ApiRouter.Xxx` 来进行定义，并且使用 `ApiRouter.Group` 函数支持多级分组。

## 路由版本及路径

如果路由的版本与 `components.ApiDefaultVersion` （即默认版本）相同，则该版本下的路由，将有两个路径可以访问，如：`/v1/user/detail/123` 和 `/user/detail/123` 在默认版本为 `v1` 时，实际上为同一路由。

该机制是为了保证在进行大版本升级时，对外展示的默认路由不必对路径中的版本号进行修改，在有大版本升级时，只需要修改默认版本号即可。但按照语义化版本规范，因为大版本号之间的 API 是不兼容的，所以实际使用时，建议使用带版本号的路径。

如果需要强行使用带版本号前缀的路由路径，只需要将 `components.ApiForcePrefixVersion` 设置为 `true`。