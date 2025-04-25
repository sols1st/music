declare module '@/components/SvgIcon.vue' {
  import { DefineComponent } from 'vue'
  const component: DefineComponent<{
    iconClass: string
    className?: string
  }>
  export default component
} 