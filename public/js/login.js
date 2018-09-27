let formLogin = $('form-login'),
        email = $('email')
        passwrod = $('password'),
        btnLogin = $('btnLogin')
        mensaje = $('mensaje');

formLogin.addEventListener('submit', e => {
    e.preventDefault();
    let obj = {
        email: email.value,
        password: password.value
    };

    peticionAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
            .then( respuesta => {
                if(respuesta.status === 200){
                    mensaje.textContent = 'ingresaste';
                    sessionStorage.setItem('token', respuesta.response.token);
                    console.log(respuesta.response);
                } else {
                    mensaje.textContent = respuesta.response.message;
                    console.log(respuesta.response);
                }
            })
            .catch(error => {
                console.log(error);
            })
})
