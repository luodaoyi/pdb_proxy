package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pdb_proxy/conf"
	"pdb_proxy/pdb"
	"strings"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	// 打印启动信息
	printStartupInfo()

	r := gin.Default()

	// 添加根路径处理器，为浏览器访问提供友好界面
	r.GET("/", handleRootAccess)

	// 原有的pdb查询路由
	r.GET("/download/symbols/:pdbname/:pdbhash/:pdbname", pdb.PdbQuery)

	// 处理其他404情况，如果是浏览器访问则返回友好页面
	r.NoRoute(handleNotFound)

	// 打印服务器启动完成信息
	fmt.Printf("🚀 服务器启动完成！访问 http://%s 获取配置说明\n", conf.ServerPort)
	fmt.Println("📝 按 Ctrl+C 停止服务器")
	fmt.Println(strings.Repeat("=", 60))

	r.Run(conf.ServerPort)
}

// printStartupInfo 打印启动信息
func printStartupInfo() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("🔧 PDB Proxy Server - Microsoft 符号服务器代理")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Printf("⏰ 启动时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("🌐 监听地址: %s\n", conf.ServerPort)
	fmt.Printf("📁 PDB 缓存目录: %s\n", conf.PdbDir)
	fmt.Printf("🔗 上游服务器: %s\n", conf.PdbServer)
	fmt.Println("📋 可用端点:")
	fmt.Println("   GET  /                                    - 配置说明页面")
	fmt.Println("   GET  /download/symbols/{name}/{hash}/{name} - PDB文件下载")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Println("�� 正在启动服务器...")
}

// handleRootAccess 处理根路径访问
func handleRootAccess(c *gin.Context) {
	// 检查是否为浏览器访问
	if isBrowserRequest(c) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, getHelpHTML())
	} else {
		c.String(http.StatusNotFound, "PDB Proxy Server - Use /download/symbols/[pdbname]/[pdbhash]/[pdbname] for PDB files")
	}
}

// handleNotFound 处理404情况
func handleNotFound(c *gin.Context) {
	// 检查是否为浏览器访问
	if isBrowserRequest(c) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusNotFound, getHelpHTML())
	} else {
		c.String(http.StatusNotFound, "pdb not exist")
	}
}

// isBrowserRequest 检查是否为浏览器请求
func isBrowserRequest(c *gin.Context) bool {
	userAgent := c.GetHeader("User-Agent")
	accept := c.GetHeader("Accept")

	// 检查User-Agent中是否包含常见浏览器标识
	browserIdentifiers := []string{"Mozilla", "Chrome", "Safari", "Firefox", "Edge", "Opera"}
	for _, identifier := range browserIdentifiers {
		if strings.Contains(userAgent, identifier) {
			return true
		}
	}

	// 检查Accept头是否包含text/html
	if strings.Contains(accept, "text/html") {
		return true
	}

	return false
}

// getHelpHTML 返回帮助页面HTML
func getHelpHTML() string {
	return `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>PDB Proxy Server - 配置说明</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            line-height: 1.6;
            margin: 0;
            padding: 20px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            background: white;
            border-radius: 10px;
            box-shadow: 0 10px 30px rgba(0,0,0,0.2);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #4CAF50 0%, #45a049 100%);
            color: white;
            padding: 30px;
            text-align: center;
        }
        .header h1 {
            margin: 0;
            font-size: 2.2em;
            font-weight: 300;
        }
        .header p {
            margin: 10px 0 0 0;
            opacity: 0.9;
            font-size: 1.1em;
        }
        .content {
            padding: 30px;
        }
        .section {
            margin-bottom: 30px;
            padding: 20px;
            background: #f8f9fa;
            border-radius: 8px;
            border-left: 4px solid #4CAF50;
        }
        .section h2 {
            color: #333;
            margin-top: 0;
            margin-bottom: 15px;
            font-size: 1.4em;
        }
        .code-block {
            background: #2d3748;
            color: #e2e8f0;
            padding: 15px;
            border-radius: 5px;
            font-family: 'Courier New', monospace;
            font-size: 0.9em;
            margin: 10px 0;
            overflow-x: auto;
            position: relative;
        }
        .copy-btn {
            position: absolute;
            top: 10px;
            right: 10px;
            background: #4CAF50;
            color: white;
            border: none;
            padding: 5px 10px;
            border-radius: 3px;
            cursor: pointer;
            font-size: 0.8em;
        }
        .copy-btn:hover {
            background: #45a049;
        }
        .variable-name {
            color: #90cdf4;
            font-weight: bold;
        }
        .alert {
            background: #fff3cd;
            border: 1px solid #ffeaa7;
            color: #856404;
            padding: 15px;
            border-radius: 5px;
            margin: 15px 0;
        }
        .success {
            background: #d4edda;
            border: 1px solid #c3e6cb;
            color: #155724;
        }
        .footer {
            text-align: center;
            padding: 20px;
            background: #f8f9fa;
            color: #666;
        }
        .footer a {
            color: #4CAF50;
            text-decoration: none;
        }
        .footer a:hover {
            text-decoration: underline;
        }
        .current-server {
            color: #e53e3e;
            font-weight: bold;
            font-family: monospace;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🔧 PDB Proxy Server</h1>
            <p>Microsoft 符号服务器代理 - 配置说明</p>
        </div>
        
        <div class="content">
            <div class="section">
                <h2>📋 环境变量配置</h2>
                <p>要使用此 PDB 代理服务器，请按以下步骤配置您的环境变量：</p>
                
                <div class="alert">
                    <strong>变量名：</strong> <span class="variable-name">_NT_SYMBOL_PATH</span>
                </div>
                
                <div class="alert success">
                    <strong>变量值：</strong>
                    <div class="code-block">
                        <button class="copy-btn" onclick="copyToClipboard(this)">复制</button>
                        <div id="symbol-path">srv*C:\Symbols*<span class="current-server"></span>/download/symbols</div>
                    </div>
                </div>
            </div>
            
            <div class="section">
                <h2>🖥️ Windows 设置步骤</h2>
                <ol>
                    <li>右键点击 <strong>此电脑</strong> → <strong>属性</strong></li>
                    <li>点击 <strong>高级系统设置</strong></li>
                    <li>点击 <strong>环境变量</strong> 按钮</li>
                    <li>在 <strong>用户变量</strong> 或 <strong>系统变量</strong> 中点击 <strong>新建</strong></li>
                    <li>变量名输入：<code>_NT_SYMBOL_PATH</code></li>
                    <li>变量值输入上面显示的路径</li>
                    <li>点击 <strong>确定</strong> 保存设置</li>
                </ol>
            </div>
            
            <div class="section">
                <h2>💡 使用说明</h2>
                <ul>
                    <li><strong>C:\Symbols</strong> - 本地符号缓存目录，可以根据需要修改</li>
                    <li><strong>srv*</strong> - 表示使用符号服务器缓存模式</li>
                    <li>配置完成后，调试器会自动从此代理服务器下载符号文件</li>
                    <li>符号文件会被缓存到本地，提高后续访问速度</li>
                </ul>
            </div>
            
            <div class="section">
                <h2>🛠️ 支持的调试器</h2>
                <ul>
                    <li>Visual Studio</li>
                    <li>WinDbg</li>
                    <li>x64dbg</li>
                    <li>OllyDbg</li>
                    <li>其他支持 Microsoft 符号服务器协议的调试器</li>
                </ul>
            </div>
        </div>
        
        <div class="footer">
            <p>更多信息访问：<a href="https://github.com/szdyg/pdb_proxy" target="_blank">https://github.com/szdyg/pdb_proxy</a></p>
            <p>当前服务器地址：<span class="current-server"></span></p>
        </div>
    </div>
    
    <script>
        // 获取当前域名和端口
        function getCurrentServer() {
            return window.location.protocol + '//' + window.location.host;
        }
        
        // 更新页面中的服务器地址
        function updateServerAddress() {
            const currentServer = getCurrentServer();
            const elements = document.querySelectorAll('.current-server');
            elements.forEach(element => {
                element.textContent = currentServer;
            });
        }
        
        // 复制到剪贴板
        function copyToClipboard(button) {
            const symbolPath = document.getElementById('symbol-path').textContent;
            
            if (navigator.clipboard) {
                navigator.clipboard.writeText(symbolPath).then(() => {
                    button.textContent = '已复制!';
                    setTimeout(() => {
                        button.textContent = '复制';
                    }, 2000);
                });
            } else {
                // 降级方案
                const textArea = document.createElement('textarea');
                textArea.value = symbolPath;
                document.body.appendChild(textArea);
                textArea.select();
                document.execCommand('copy');
                document.body.removeChild(textArea);
                
                button.textContent = '已复制!';
                setTimeout(() => {
                    button.textContent = '复制';
                }, 2000);
            }
        }
        
        // 页面加载时更新服务器地址
        document.addEventListener('DOMContentLoaded', updateServerAddress);
    </script>
</body>
</html>`
}
