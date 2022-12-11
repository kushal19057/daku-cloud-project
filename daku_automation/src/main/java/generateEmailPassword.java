import java.util.ArrayList;

public class generateEmailPassword {

    public static void main(String[] args) {
        int n=50;



        ArrayList<String> emails=new ArrayList<String>(n);
        ArrayList<String> passwords=new ArrayList<String>(n);

        for(int i=0;i<n;i++)
        {
//            System.out.println(i);
            emails.add("email"+Integer.toString(i)+"@gmail.com");
            passwords.add("password"+Integer.toString(i));
        }

        for(int i=0;i<n;i++)
        {
//            System.out.println(i);
            System.out.println(emails.get(i));
            System.out.println(passwords.get(i));
        }




    }
}
