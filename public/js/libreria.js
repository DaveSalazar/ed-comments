function peticionAjax(metodo, url, obj) {
    return new Promise(function(resolver,rechazar) {
        let xhr = new XMLHttpRequest();

        xhr.open(metodo, url, true);
        xhr.setRequestHeader('Content-Type','application/json');
        xhr.addEventListener('load', e => {
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            }

            resolver(respuesta);
        });
        
        xhr.addEventListener('error', e => {
            let self = e.target;
            rechazar(self);
        });
        
        xhr.send(obj);
    });
}

function $(elemento){
    return document.getElementById(elemento);
}