

function mostrar(){
    return new Promise((resolve,reject)=>{
        console.log(document.getElementById('ho'))
        
        escribe()
        resolve(escribe())
    }

    );
}

function escribe(){
    setTimeout(()=>{
        for (let i=0;i<5;i++){
            document.getElementById('lista').innerHTML="<h1>Parrafo</h1>";
    
        }
    },100)
}


escribe()