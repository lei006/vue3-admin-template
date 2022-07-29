
const Router = require('koa-router');
const userCtl = require('../controller/user.controller');
const authCtl = require('../controller/auth.controller');

const jwt = require('koa-jwt');


const router = new Router({prefix: '/api'});

// 指定一个url匹配
router.get('/', async (ctx) => {
    ctx.type = 'html';
    ctx.body = '<h1>hello world!</h1>';
})

console.log("process.env.JWT_SECRET", process.env.JWT_SECRET);
router.use(jwt({secret:process.env.JWT_SECRET}).unless({
    path:[/^\/api\/auth\/login/,/^\/api\/auth\/captcha/]
}));

/*
//////////////////////////////////////////////
// 验证 jwt
const secret = 'moyufed-test';
router.use(jwt({
    secret,
    debug: true // 开启debug可以看到准确的错误信息
    })
    .unless({ path: [/\/register/, /\/login/] }) // 以 public 开头的请求地址不使用 jwt 中间件
);
*/


const authRouter = new Router();
authRouter.post('/login', authCtl.login);
authRouter.post('/logout', authCtl.logout);
authRouter.get('/info', authCtl.info);
authRouter.post('/captcha', authCtl.captcha);

router.use('/auth', authRouter.routes(), authRouter.allowedMethods())


const userRouter = new Router();
userRouter.post('/', userCtl.add);
userRouter.put('/:id', userCtl.update);
userRouter.get('/', userCtl.list);
userRouter.delete('/:id', userCtl.delete);
userRouter.get('/:id', userCtl.find);

router.use('/user', userRouter.routes(), userRouter.allowedMethods())


//router.use('/user', user.routes(), user.allowedMethods());

module.exports = router;

