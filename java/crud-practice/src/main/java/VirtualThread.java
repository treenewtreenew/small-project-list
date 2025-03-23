import java.time.Duration;
import java.time.Instant;

public class VirtualThread {
    private final static Duration EMULATE_WORK_FOR = Duration.ofSeconds(5);

    public static void _main(String[] args) throws InterruptedException{
        Instant start = Instant.now();

        int maxThreadNo = 1_000;
        Thread[] threads = new Thread[maxThreadNo + 1];
        for (int i = 0; i < maxThreadNo; i++) {
            threads[i] = createThread(i);
            threads[i].start();
            if (i % 5_000 == 0) {
                System.out.printf("Current count %d%n", i);
            }
        }
    }

    private static Thread createThread(int i) {
        Runnable job = () -> blockingOperaton(i);
        return new Thread(job);
    }

    static void blockingOperaton(int task) {
        sneakySleep(EMULATE_WORK_FOR);
    }

    public static void sneakySleep(Duration d) {
        try {
            Thread.sleep(1L);
        } catch (InterruptedException e){
            throw new RuntimeException(e);
        }
    }
}
