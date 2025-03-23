import java.time.Instant;
import java.util.Hashtable;

public class HashTableTest {
    public static void main(String[] args) throws InterruptedException{
        Hashtable<Integer, String> table = new Hashtable<>();
        table.put(1, "One");
        table.put(2, "Two");
        table.put(3, "Three");
        table.put(4, "Four");

        table.computeIfAbsent(5, k -> "Five");
        table.computeIfPresent(2, (k, v) -> null);
        table.computeIfAbsent(3, k -> "ThreeNew");

        System.out.println(table);
    }
}
