import './iconfont.js'
import './iconfont1.js'
import './iconfont2.js'
import './iconfont3.js'
import './iconfont4.js'
import SvgIcon from '@/components/SvgIcon'

// 全局注册组件
export function setupSvgIcon(app) {
  app.component('svg-icon', SvgIcon)
  
  // 导入所有的 svg 图标
  const requireAll = requireContext => requireContext.keys().map(requireContext)
  const req = require.context('./', true, /\.svg$/)
  requireAll(req)
}