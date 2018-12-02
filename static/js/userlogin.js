$(function() {
    //获取焦点事件
    $('#callnum').focus(function() {
        //获取当前元素的值
        this.value = (this.value == '手机号/邮箱') ? '' : this.value;
    });
    $('#callpass').focus(function() {
        //获取当前元素的值
        this.value = (this.value == '******') ? '' : this.value;
    });
    //绑定按钮事件
    // $('#btn_login').click(function(e) {
    //     //输出
    //     console.log(e);
    //     //获取用户名和密码
    //     var name = $('#callnum').value();
    //     var password = md5($('#callpass').value());
    //     //密码一般用md5码进行加密

    //     //提示框
    //     window.wxc.xcConfirm(name + '<br/>' + password, window.wxc.xcConfirm.typeEnum.success);

    //     setTimeout(function() {
    //         location.href = 'index.html';
    //     }, 3000);
    // });
    $('#telphone').focus(function() {
        //获取当前元素的值
        this.value = (this.value == '手机号/邮箱') ? '' : this.value;
    });
    $('#password').focus(function() {
        //获取当前元素的值
        this.value = (this.value == '******') ? '' : this.value;
    });
})