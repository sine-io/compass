本项目的主干。

每个应用程序的目录名应该与你想要的可执行文件的名称相匹配(例如，/cmd/myapp)。

不要在这个目录中放置太多代码。如果你认为代码可以导入并在其他项目中使用，那么它应该位于 /pkg 目录中。
如果代码不是可重用的，或者你不希望其他人重用它，请将该代码放到 /internal 目录中。
你会惊讶于别人会怎么做，所以要明确你的意图!

通常有一个小的 main 函数，从 /internal 和 /pkg 目录导入和调用代码，除此之外没有别的东西。

有关示例，请参阅 /cmd 目录。

Examples:
https://github.com/vmware-tanzu/velero/tree/main/cmd (just a really small main function with everything else in packages)
https://github.com/moby/moby/tree/master/cmd
https://github.com/prometheus/prometheus/tree/main/cmd
https://github.com/influxdata/influxdb/tree/master/cmd
https://github.com/kubernetes/kubernetes/tree/master/cmd
https://github.com/dapr/dapr/tree/master/cmd
https://github.com/ethereum/go-ethereum/tree/master/cmd
