/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package escuela;

import com.google.gson.Gson;
import javax.swing.table.DefaultTableModel;

/**
 *
 * @author scion833
 */
public class frameProfesor extends javax.swing.JFrame {

    /**
     * Creates new form frameProfesor
     */
    public frameProfesor() {
        initComponents();
        leerJson();
        llenarTabla();
    }

    private Profesor[] profesores;

    public void leerJson() {
        consumiendoApi ca = new consumiendoApi();
        String json = ca.jsonGetRequest("http://localhost:8000/api/v1/profesores/");
        Gson objGson = new Gson();
        respuesta res = new respuesta();
        res = objGson.fromJson(json, respuesta.class);
        this.profesores = res.data;
    }

    public void llenarTabla() {
        leerJson();
        resetearTabla();
        for(int i=0;i<profesores.length;i++){
            agregarFilaTabla();
            this.jTable1.setValueAt(profesores[i].ci+"", i, 0);
            this.jTable1.setValueAt(profesores[i].nombres, i, 1);
            this.jTable1.setValueAt(profesores[i].paterno, i, 2);
            this.jTable1.setValueAt(profesores[i].materno, i, 3);
            this.jTable1.setValueAt(profesores[i].correo, i, 4);
        }
    }
    
    public void agregarFilaTabla(){
        DefaultTableModel model=(DefaultTableModel) this.jTable1.getModel();
        model.addRow(new Object[] {""});
    }

    public void resetearTabla() {
        jTable1.setModel(new javax.swing.table.DefaultTableModel(
                new Object[][]{},
                new String[]{
                    "Ci", "Nombres", "Paterno", "Materno", "Correo"
                }
        ));
    }

    /**
     * This method is called from within the constructor to initialize the form.
     * WARNING: Do NOT modify this code. The content of this method is always
     * regenerated by the Form Editor.
     */
    @SuppressWarnings("unchecked")
    // <editor-fold defaultstate="collapsed" desc="Generated Code">//GEN-BEGIN:initComponents
    private void initComponents() {

        jLabel2 = new javax.swing.JLabel();
        jButton1 = new javax.swing.JButton();
        jScrollPane1 = new javax.swing.JScrollPane();
        jTable1 = new javax.swing.JTable();
        pantCi = new javax.swing.JTextField();
        jLabel1 = new javax.swing.JLabel();
        jLabel3 = new javax.swing.JLabel();
        pantNombres = new javax.swing.JTextField();
        jLabel4 = new javax.swing.JLabel();
        pantMaterno = new javax.swing.JTextField();
        jLabel5 = new javax.swing.JLabel();
        pantPaterno = new javax.swing.JTextField();
        jLabel6 = new javax.swing.JLabel();
        jLabel7 = new javax.swing.JLabel();
        pantCorreo = new javax.swing.JTextField();
        pantPassword = new javax.swing.JPasswordField();
        jButton2 = new javax.swing.JButton();
        jButton3 = new javax.swing.JButton();
        jButton4 = new javax.swing.JButton();
        jLabel8 = new javax.swing.JLabel();

        setDefaultCloseOperation(javax.swing.WindowConstants.EXIT_ON_CLOSE);
        getContentPane().setLayout(new org.netbeans.lib.awtextra.AbsoluteLayout());

        jLabel2.setFont(new java.awt.Font("Dialog", 3, 18)); // NOI18N
        jLabel2.setHorizontalAlignment(javax.swing.SwingConstants.CENTER);
        jLabel2.setText("PROFESORES");
        getContentPane().add(jLabel2, new org.netbeans.lib.awtextra.AbsoluteConstraints(220, 10, 170, 50));

        jButton1.setText("ACTUALIZAR LISTA");
        jButton1.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                jButton1ActionPerformed(evt);
            }
        });
        getContentPane().add(jButton1, new org.netbeans.lib.awtextra.AbsoluteConstraints(190, 270, 210, -1));

        jTable1.setModel(new javax.swing.table.DefaultTableModel(
            new Object [][] {
            },
            new String [] {
                "Ci", "Nombres", "Paterno", "Materno","Correo"
            }
        ));
        jScrollPane1.setViewportView(jTable1);

        getContentPane().add(jScrollPane1, new org.netbeans.lib.awtextra.AbsoluteConstraints(20, 310, 550, 150));

        pantCi.addActionListener(new java.awt.event.ActionListener() {
            public void actionPerformed(java.awt.event.ActionEvent evt) {
                pantCiActionPerformed(evt);
            }
        });
        getContentPane().add(pantCi, new org.netbeans.lib.awtextra.AbsoluteConstraints(20, 80, 120, -1));

        jLabel1.setText("Ci:");
        getContentPane().add(jLabel1, new org.netbeans.lib.awtextra.AbsoluteConstraints(20, 60, 80, 20));

        jLabel3.setText("Ingrese Contraseña:");
        getContentPane().add(jLabel3, new org.netbeans.lib.awtextra.AbsoluteConstraints(300, 120, 180, 20));
        getContentPane().add(pantNombres, new org.netbeans.lib.awtextra.AbsoluteConstraints(160, 80, 110, -1));

        jLabel4.setText("Nombres:");
        getContentPane().add(jLabel4, new org.netbeans.lib.awtextra.AbsoluteConstraints(160, 60, -1, 20));
        getContentPane().add(pantMaterno, new org.netbeans.lib.awtextra.AbsoluteConstraints(450, 80, 120, -1));

        jLabel5.setText("Ap. Paterno:");
        getContentPane().add(jLabel5, new org.netbeans.lib.awtextra.AbsoluteConstraints(300, 60, -1, 20));
        getContentPane().add(pantPaterno, new org.netbeans.lib.awtextra.AbsoluteConstraints(300, 80, 120, -1));

        jLabel6.setText("Ap. Materno:");
        getContentPane().add(jLabel6, new org.netbeans.lib.awtextra.AbsoluteConstraints(450, 60, -1, 20));

        jLabel7.setText("Ingrese Correo:");
        getContentPane().add(jLabel7, new org.netbeans.lib.awtextra.AbsoluteConstraints(120, 120, 140, 20));
        getContentPane().add(pantCorreo, new org.netbeans.lib.awtextra.AbsoluteConstraints(120, 140, 140, -1));

        pantPassword.setMinimumSize(new java.awt.Dimension(15, 23));
        pantPassword.setPreferredSize(new java.awt.Dimension(15, 23));
        getContentPane().add(pantPassword, new org.netbeans.lib.awtextra.AbsoluteConstraints(300, 140, 170, -1));

        jButton2.setText("ACEPTAR");
        getContentPane().add(jButton2, new org.netbeans.lib.awtextra.AbsoluteConstraints(250, 180, -1, -1));

        jButton3.setText("GUARDAR CAMBIOS");
        getContentPane().add(jButton3, new org.netbeans.lib.awtextra.AbsoluteConstraints(220, 220, -1, -1));

        jButton4.setText("ELIMINAR");
        getContentPane().add(jButton4, new org.netbeans.lib.awtextra.AbsoluteConstraints(460, 270, -1, -1));
        getContentPane().add(jLabel8, new org.netbeans.lib.awtextra.AbsoluteConstraints(0, 0, 590, 480));

        pack();
    }// </editor-fold>//GEN-END:initComponents

    private void jButton1ActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_jButton1ActionPerformed
        // TODO add your handling code here:
    }//GEN-LAST:event_jButton1ActionPerformed

    private void pantCiActionPerformed(java.awt.event.ActionEvent evt) {//GEN-FIRST:event_pantCiActionPerformed
        // TODO add your handling code here:
    }//GEN-LAST:event_pantCiActionPerformed

    /**
     * @param args the command line arguments
     */
    public static void main(String args[]) {
        /* Set the Nimbus look and feel */
        //<editor-fold defaultstate="collapsed" desc=" Look and feel setting code (optional) ">
        /* If Nimbus (introduced in Java SE 6) is not available, stay with the default look and feel.
         * For details see http://download.oracle.com/javase/tutorial/uiswing/lookandfeel/plaf.html 
         */
        try {
            for (javax.swing.UIManager.LookAndFeelInfo info : javax.swing.UIManager.getInstalledLookAndFeels()) {
                if ("Nimbus".equals(info.getName())) {
                    javax.swing.UIManager.setLookAndFeel(info.getClassName());
                    break;
                }
            }
        } catch (ClassNotFoundException ex) {
            java.util.logging.Logger.getLogger(frameProfesor.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (InstantiationException ex) {
            java.util.logging.Logger.getLogger(frameProfesor.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (IllegalAccessException ex) {
            java.util.logging.Logger.getLogger(frameProfesor.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        } catch (javax.swing.UnsupportedLookAndFeelException ex) {
            java.util.logging.Logger.getLogger(frameProfesor.class.getName()).log(java.util.logging.Level.SEVERE, null, ex);
        }
        //</editor-fold>

        /* Create and display the form */
        java.awt.EventQueue.invokeLater(new Runnable() {
            public void run() {
                new frameProfesor().setVisible(true);
            }
        });
    }

    // Variables declaration - do not modify//GEN-BEGIN:variables
    private javax.swing.JButton jButton1;
    private javax.swing.JButton jButton2;
    private javax.swing.JButton jButton3;
    private javax.swing.JButton jButton4;
    private javax.swing.JLabel jLabel1;
    private javax.swing.JLabel jLabel2;
    private javax.swing.JLabel jLabel3;
    private javax.swing.JLabel jLabel4;
    private javax.swing.JLabel jLabel5;
    private javax.swing.JLabel jLabel6;
    private javax.swing.JLabel jLabel7;
    private javax.swing.JLabel jLabel8;
    private javax.swing.JScrollPane jScrollPane1;
    private javax.swing.JTable jTable1;
    private javax.swing.JTextField pantCi;
    private javax.swing.JTextField pantCorreo;
    private javax.swing.JTextField pantMaterno;
    private javax.swing.JTextField pantNombres;
    private javax.swing.JPasswordField pantPassword;
    private javax.swing.JTextField pantPaterno;
    // End of variables declaration//GEN-END:variables
}
