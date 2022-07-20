package az.ingress.statemachine.account;

import az.ingress.statemachine.account.transitions.Approve;
import az.ingress.statemachine.account.transitions.Notify;
import az.ingress.statemachine.account.transitions.Reject;
import az.ingress.statemachine.account.transitions.Submit;

import java.util.Arrays;
import java.util.List;

public enum AccountStatus {

    DRAFT(Submit.NAME),
    IN_REVIEW(Approve.NAME, Reject.NAME),
    APPROVED(Approve.NAME, Notify.NAME),
    NOTIFIED();

    private final List<String> transitions;

    AccountStatus(String... transitions) {
        this.transitions = Arrays.asList(transitions);
    }

    public List<String> getTransitions() {
        return transitions;
    }
}
