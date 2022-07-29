
class UserController{
    async index() {
      const { ctx } = this;
      ctx.body = 'hi, user';
    }
}


module.exports = UserController;
