

function mostrar(){
    return new Promise((resolve,reject)=>{
        console.log(document.getElementById('ho'))
        
        escribe()
        resolve(escribe())
    }

    );
}

function escribe(){
    //for (let i=0;i<5;i++){
        document.getElementById('lista').innerHTML="Parrafo";

    //}
}

mostrar()
.then(function(){
    escribe()
}).catch(function(err){
    console.log(err)
});