package az.ingress.statemachine;

import az.ingress.statemachine.account.AccountDto;
import az.ingress.statemachine.account.AccountStatus;
import az.ingress.statemachine.account.TransitionService;
import az.ingress.statemachine.account.transitions.Approve;
import az.ingress.statemachine.account.transitions.Notify;
import az.ingress.statemachine.account.transitions.Reject;
import az.ingress.statemachine.account.transitions.Submit;
import az.ingress.statemachine.domain.Account;
import az.ingress.statemachine.repository.AccountRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * https://medium.com/@iyusubov444 Implementing a Custom State Machine with Spring Boot
 */
@SpringBootApplication
@RequiredArgsConstructor
public class StateMachineApplication implements CommandLineRunner {

    private final TransitionService<AccountDto> accountTransitionService;
    private final AccountRepository accountRepository;

    public static void main(String[] args) {
        SpringApplication.run(StateMachineApplication.class, args);
    }

    @Override
    public void run(String... args) throws Exception {
        Account account = new Account();
        account.setStatus(AccountStatus.DRAFT);
        accountRepository.save(account);
        accountTransitionService.getAllowedTransitions(account.getId())
                .stream()
                .forEach(System.out::println);

        //Draft -> Submitted
        accountTransitionService.transition(account.getId(), Submit.NAME);

        //IN_REVIEW
        accountTransitionService.transition(account.getId(), Reject.NAME);

        //Draft -> Submitted
        accountTransitionService.transition(account.getId(), Submit.NAME);

        //InReview -> Approved
        accountTransitionService.transition(account.getId(), Approve.NAME);

        //Approve -> Approve
        accountTransitionService.transition(account.getId(), Approve.NAME);

        //Approve -> Approve
        accountTransitionService.transition(account.getId(), Approve.NAME);

        //Approve -> Approve
        accountTransitionService.transition(account.getId(), Approve.NAME);

        //Approve -> Approve
        accountTransitionService.transition(account.getId(), Approve.NAME);

        //Approve -> Notify
        accountTransitionService.transition(account.getId(), Notify.NAME);
    }

}
