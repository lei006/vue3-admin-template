import router from './router'


router.beforeEach(async(to, from, next) => {
    console.log('beforeEach', to, from)
    next()
})

router.afterEach(() => {
  // finish progress bar
  console.log('afterEach')
  //NProgress.done()
})
