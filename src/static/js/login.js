
jQuery(document).ready(function() {
	
    /*
        Fullscreen background
    */
	$('#zhuce').click(function(){
		$.post('YongHu_zhuce',
				{denglm : $('#denglm').val() ,
				mim : $('#mim').val() , 
				bh : $('#bh').val() ,
				xingm : $('#xingm').val()
				},
				function(data){
					if(data = {"result" : 1}){alert('注册成功，你现在可以使用此账号登录');}
					else if(data = {"result" : 0}){alert('系统繁忙，注册失败');}
					else{alert('出现为之错误');}
				});
	});
    $.backstretch("/static/images/loginbackground.jpg");
    
    /*
        Form validation
    */
    $('.login-form input[type="text"], .login-form input[type="password"], .login-form textarea').on('focus', function() {
    	$(this).removeClass('input-error');
    });
    
    $('.login-form').on('submit', function(e) {
    	
    	$(this).find('input[type="text"], input[type="password"], textarea').each(function(){
    		if( $(this).val() == "" ) {
    			e.preventDefault();
    			$(this).addClass('input-error');
    		}
    		else {
    			$(this).removeClass('input-error');
    		}
    	});
    	
    });
    
    
});