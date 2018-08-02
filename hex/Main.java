public class Main {
    public static String byteArrayToHex(byte[] a) {
        StringBuilder sb = new StringBuilder(a.length * 2);
        for(byte b: a)
            sb.append(String.format("%02x", b));
        return sb.toString();
    }

    public static void main(String args[]) throws Exception {
        System.out.println("hello");
        String str = "Hello World";
        String ret = byteArrayToHex(str.getBytes("UTF-8"));
        System.out.println(ret);
    }
}
