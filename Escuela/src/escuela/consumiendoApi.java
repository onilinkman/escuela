
package escuela;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import java.util.Scanner;


public class consumiendoApi {
    
    private String metodo="GET";
    private String llave="Content-Type";
    private String valor="application/json";
    private String json;
    
    public consumiendoApi(){
        
    }
    
    public consumiendoApi(String json,String metodo){
        this.json=json;
        this.metodo=metodo;
    }
    
    
    private String streamToString(InputStream inputStream) throws IOException{
        String text=new Scanner(inputStream,"UTF-8").useDelimiter("\\z").next();
        return text;
    }
    
    
    public String jsonGetRequest(String urlQueryString){
        String json=null;
        try{
            URL url=new URL(urlQueryString);
            HttpURLConnection connection=(HttpURLConnection) url.openConnection();
            connection.setDoOutput(true);
            connection.setInstanceFollowRedirects(false);
            connection.setRequestMethod(metodo);
            connection.setRequestProperty(llave, valor);
            connection.setRequestProperty("charset", "utf-8");
            connection.connect();
            InputStream inStream=connection.getInputStream();
            json=streamToString(inStream);
            
        }catch(Exception ex){
            System.out.println(ex);
        }
        return json;
    }
    
    public void postRequest(String dir){
        try{
            URL url=new URL(dir);
            HttpURLConnection con=(HttpURLConnection) url.openConnection();
            con.setRequestMethod(metodo);
            con.setRequestProperty(llave, "application/json; utf-8");
            con.setRequestProperty("Accept", valor);
            con.setDoOutput(true);
            
            try(OutputStream os=con.getOutputStream()){
                byte[] input=this.json.getBytes("utf-8");
                os.write(input, 0, input.length);
            }
            try (BufferedReader br=new BufferedReader(new InputStreamReader(con.getInputStream(),"utf-8"))){
                StringBuilder response=new StringBuilder();
                String responseLine=null;
                while((responseLine=br.readLine())!=null){
                    response.append(responseLine.trim());
                }
            }
            
        }catch(Exception e){
            System.out.println("Error: "+e);
        }
    }
    
    
    
}
