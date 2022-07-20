package az.ingress.statemachine.account;

public interface Transition<T> {

    String getName();

    AccountStatus getTargetStatus();

    //This should do required pre processing
    void applyProcessing(T order);
}
