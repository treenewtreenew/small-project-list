package az.ingress.statemachine.repository;

import az.ingress.statemachine.domain.Account;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AccountRepository extends JpaRepository<Account, Long> {
}
