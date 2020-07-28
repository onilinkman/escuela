/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package escuela;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;
import java.lang.reflect.Type;
import java.util.List;

/**
 *
 * @author scion833
 */
public class Escuela {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        Gson objGson=new Gson();
        classUE ue=new classUE();
        ue.nombre="unidad educativa";
        ue.direccion="san francisco Nº456";
        ue.correo="sanfrancisco@gmail.com";
        String jsonue=objGson.toJson(ue);
        //System.out.println(jsonue);
        //consumiendoApi csa=new consumiendoApi(jsonue, "POST");
        consumiendoApi csa=new consumiendoApi();
        //csa.postRequest("http://localhost:8000/api/v1/ue");
        String json=csa.jsonGetRequest("http://localhost:8000/api/v1/profesores/");
        respuesta res=new respuesta();
        res=objGson.fromJson(json, respuesta.class);
        //Type tipoListaProfes= new TypeToken<List<Profesor>>(){}.getType();
        //List<Profesor> profes=objGson.fromJson(res.data., tipoListaProfes);
        Profesor[] obje=res.data;
        
        
        System.out.println(obje.length);
        
    }
    
}
