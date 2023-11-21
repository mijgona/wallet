using DataAccess;

namespace BusinessLogic;

public interface ITransactionService
{
    public ValueTask<Transaction> CreateTransactionAsync(TransactionInfo transaction, CancellationToken token);

    public ValueTask<List<Transaction>> GetTransactionsAsync(long userId, CancellationToken token);
}
