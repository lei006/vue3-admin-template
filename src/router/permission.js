import NProgress from '@/utils/nprogress'



let Permission = (router)=>{
    router.beforeEach((to, from, next) => {
        console.log("before: ",to)
        next()
    })

    router.afterEach(() => {
        // finish progress bar
        console.log("after: ")
        NProgress.done()
    })
}


export default Permission

