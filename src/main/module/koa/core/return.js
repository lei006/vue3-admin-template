"use strict"


class Return {

    async ok(ctx, data){
      ctx.body = {code:20000, data, msg:"ok"};
    }
    async success(ctx, data){
      ctx.body = {code:20000, data, msg:"ok"};
    }
    
    async error(ctx, code, msg){
      ctx.body = {code, msg};
    }

    async _return(code, data){
      let ret = {code, data};
      return ret;
    }

}

module.exports = new Return();

