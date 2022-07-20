package az.ingress.statemachine.handler;

import az.ingress.statemachine.events.AccountTransitionedEvent;
import lombok.extern.slf4j.Slf4j;
import org.springframework.context.ApplicationListener;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class AccountSubmittedEventHandler implements ApplicationListener<AccountTransitionedEvent> {

    @Override
    public void onApplicationEvent(AccountTransitionedEvent event) {
        log.info("Received account status changed event for account  {}  from state {} to state {} from source{}",
                event.getAccountDto().getId(), event.getAccountStatus(),
                event.getAccountDto().getStatus(), event.getSource());
    }
}
