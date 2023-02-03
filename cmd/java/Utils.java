public class Utils {
    public static volatile utils INSTANCE;

    private Utils(){}

    public static Utils getInstance(){
        if (INSTANCE==null){
            synchronized (Utils.class){
                if (INSTANCE==null){
                    INSTANCE=new Utils();
                }
            }
        }
        return INSTANCE;
    }
}