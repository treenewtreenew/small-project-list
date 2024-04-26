package com.spring3react.crudpractice.practice;

import sun.jvm.hotspot.utilities.Assert;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class CompletableFutureLearn {
    // Using CompletableFuture as a Simple Future
    public Future<String> calculateAsync() throws InterruptedException {
        CompletableFuture<String> completableFuture = new CompletableFuture<>();
        Executors.newCachedThreadPool().submit(() -> {
            Thread.sleep(800);
            completableFuture.complete("Hello");
            return null;
        });
        return completableFuture;
    }

    // CompletableFuture with Encapsulated Computation Logic
    public void asyncLogic() {
        CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> "Hello");
    }

    // Processing Results of Asynchronous Computations
    public void processResultsOfAsynchronousComputations() throws Exception{
        CompletableFuture<String> completableFuture = CompletableFuture.supplyAsync(() -> "Hello");
        CompletableFuture<String> future1 = completableFuture.thenApply(s -> s + " world");
        CompletableFuture<Void> future2 = completableFuture.thenAccept(s -> System.out.println("Computation returned: " + s));
        CompletableFuture<Void> future3 = completableFuture.thenRun(() -> System.out.println("Computation finished."));
    }

    // Combining Futures
    public void combiningFutures() {
        CompletableFuture<String> completableFuture1 = CompletableFuture.supplyAsync(() -> "Hello")
                .thenCompose(s -> CompletableFuture.supplyAsync(() -> s + " World"));
        CompletableFuture<String> completableFuture2 = CompletableFuture.supplyAsync(() -> "Hello")
                .thenCombine(CompletableFuture.supplyAsync(() -> " World"), (s1, s2) -> s1 + s2);
    }

    // Running Multiple Futures in Parallel
    public void multipleFuture() throws ExecutionException, InterruptedException {
        CompletableFuture<String> future1
                = CompletableFuture.supplyAsync(() -> "Hello");
        CompletableFuture<String> future2
                = CompletableFuture.supplyAsync(() -> "Beautiful");
        CompletableFuture<String> future3
                = CompletableFuture.supplyAsync(() -> "World");

        CompletableFuture<Void> combinedFuture
                = CompletableFuture.allOf(future1, future2, future3);

        combinedFuture.get();

        String combined = Stream.of(future1, future2, future3)
                .map(CompletableFuture::join)
                .collect(Collectors.joining(" "));
    }
}
