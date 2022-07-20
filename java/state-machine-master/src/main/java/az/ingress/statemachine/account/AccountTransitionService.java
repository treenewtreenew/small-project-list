package az.ingress.statemachine.account;

import az.ingress.statemachine.domain.Account;
import az.ingress.statemachine.events.AccountTransitionedEvent;
import az.ingress.statemachine.repository.AccountRepository;
import lombok.RequiredArgsConstructor;
import lombok.var;
import org.modelmapper.ModelMapper;
import org.springframework.context.ApplicationEventPublisher;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.PostConstruct;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;

@Service
@RequiredArgsConstructor
public class AccountTransitionService implements TransitionService<AccountDto> {

    private Map<String, Transition> transitionsMap;
    private final List<Transition> transitions;
    private final ModelMapper mapper;
    private final ApplicationEventPublisher applicationEventPublisher;
    private final AccountRepository accountRepository;


    /**
     * Returns a list of transitions allowed for a particular order identified by the id
     *
     * @param id: the id  of the order
     * @return: list of  transitions allowed
     */
    public List<String> getAllowedTransitions(Long id) {
        Account account = accountRepository.findById(id)
                .orElseThrow(() -> new IllegalArgumentException("Unknown account: " + id));
        return account.getStatus().getTransitions();
    }

    /**
     * Transitions the order from the current state to the target state
     *
     * @param id:         the id of the account
     * @param transition: the status to transition to
     * @return: the order details
     */
    @Transactional
    public AccountDto transition(Long id, String transition) {
        Transition accountTransition = transitionsMap.get(transition.toLowerCase());
        if (accountTransition == null) {
            throw new IllegalArgumentException("Unknown transition: " + transition);
        }
        return accountRepository.findById(id)
                .map(order -> {
                    checkAllowed(accountTransition, order.getStatus());
                    accountTransition.applyProcessing(mapper.map(order, AccountDto.class));
                    return updateStatus(order, accountTransition.getTargetStatus());
                })
                .map(u -> mapper.map(u, AccountDto.class))
                .orElseThrow(() -> new IllegalArgumentException("Unknown order: " + id));
    }

    @PostConstruct
    private void initTransitions() {
        Map<String, Transition> transitionHashMap = new HashMap<>();
        for (Transition accountTransition : transitions) {
            if (transitionHashMap.containsKey(accountTransition.getName())) {
                throw new IllegalStateException("Duplicate transition :" + accountTransition.getName());
            }
            transitionHashMap.put(accountTransition.getName(), accountTransition);
        }
        this.transitionsMap = transitionHashMap;
    }

    private Object updateStatus(Account order, AccountStatus targetStatus) {
        AccountStatus existingStatus = order.getStatus();
        order.setStatus(targetStatus);
        Account updated = accountRepository.save(order);

        var event = new AccountTransitionedEvent(this, existingStatus, mapper.map(updated, AccountDto.class));
        applicationEventPublisher.publishEvent(event);
        return updated;
    }

    private void checkAllowed(Transition accountTransition, AccountStatus status) {
        Set<AccountStatus> allowedSourceStatuses = Stream.of(AccountStatus.values())
                .filter(s -> s.getTransitions().contains(accountTransition.getName()))
                .collect(Collectors.toSet());
        if (!allowedSourceStatuses.contains(status)) {
            throw new RuntimeException("The transition from the current state " + status.name() + "  to the target state "
                    + accountTransition.getTargetStatus().name() + "  is not allowed!");
        }
    }

}
