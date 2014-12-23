import java.io.*;
public class Parser {

	public static void main(String args[]) throws Exception {
		FileReader fr = new FileReader("WoWCombatLog.txt");
		BufferedReader br = new BufferedReader(fr);
		FileWriter fw = null;
		String s;
		String fname = null;
		while((s = br.readLine()) != null){
			String linedate = s.substring(0, s.indexOf(' ')).replace('/', '_');
			if( linedate.equals(fname) == false){
				fname = linedate;
				if(fw != null){
					fw.flush();
					fw.close();
				} 
				fw = new FileWriter("output/" + fname + ".txt");
                System.out.println("New date: " + fname);
			}
			fw.write(s + "\n");			
		}
		fw.flush();
		fw.close();
	}
}
