私有应用程序和库代码。
这是你不希望其他人在其应用程序或库中导入代码。
请注意，这个布局模式是由 Go 编译器本身执行的。
有关更多细节，请参阅Go 1.4 release notes 。
注意，你并不局限于顶级 internal 目录。
在项目树的任何级别上都可以有多个内部目录。

你可以选择向 internal 包中添加一些额外的结构，以分隔共享和非共享的内部代码。
这不是必需的(特别是对于较小的项目)，但是最好有有可视化的线索来显示预期的包的用途。
你的实际应用程序代码可以放在 /internal/app 目录下(例如 /internal/app/myapp)，
这些应用程序共享的代码可以放在 /internal/pkg 目录下(例如 /internal/pkg/myprivlib)。

Examples:
https://github.com/hashicorp/terraform/tree/main/internal
https://github.com/influxdata/influxdb/tree/master/internal
https://github.com/perkeep/perkeep/tree/master/internal
https://github.com/jaegertracing/jaeger/tree/main/internal
https://github.com/moby/moby/tree/master/internal
https://github.com/satellity/satellity/tree/main/internal

/internal/pkg Examples:
https://github.com/hashicorp/waypoint/tree/main/internal/pkg