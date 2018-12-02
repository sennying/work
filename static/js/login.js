$(function () {
    //获取焦点事件
    $('#UserName').focus(function () {
        //获取当前的元素的值
        this.value = (this.value == '用户名/UID') ? '' : this.value;
    });
    $('#passWord').focus(function () {
        this.value = (this.value == '******') ? '' : this.value;
    });
    $('#btn_login').click(function (e) {
        //输出
        console.log(e);
        //获取输入的用户名和密码
        var name = $('#UserName').val();
        var passWord = md5($('#passWord').val());
        //密码一般都使用md5加密方式进行加密
        //或者自行定义加密解密方式进行验证登录
        //提示框
        window.wxc.xcConfirm(name + '<br/>' + password, window.wxc.xcConfirm.typeEnum.success);
        //setTimeout等多少秒才执行
        // setTimeout(() => {
        //     //页面跳转--最直接的跳转，不带任何参数或者变量
        //     location.href='index.html';
        // }, 3000);
    });


    /*
    //一下代码没有实际意义，仅作为一个案列
    //登录操作ajax异步请求
    $("").click(()=>{
        var naem=$("#UserName").val;
        var password=md5($("passWord").val());
        $.ajax({
            type:"post",
            url:"/login",
            verify:false, //登录不需要验证
            data:{
                num:nane,
                pass:password.substr(1,30)
            },
            succ:res=>{
                //箭头函数
                if(res.code==200){
                    //保存登录信息
                    util.setToken(res.data);
                    //获取当前用户的权限
                    var auth = util.getAuth();
                    if(auth.isStu){
                        //是学生
                        //updateUserInfo():
                        location.href="/user.html";
                    }else{
                        var msg="请在合适的地方登录";
                        window.wxc.xcConfirm(msg,window.wxc.xcConfirm.typeEnum.console.error);
                    }
                }else{
                    window.wxc.xcConfirm(res.msg,window.wxc.xcConfirm.typeEnum.error);
                }
            }
        });
    });*/
});
function name(params) {

}

