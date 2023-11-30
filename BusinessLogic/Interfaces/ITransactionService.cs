using DataAccess;

namespace BusinessLogic;

public interface ITransactionService
{
    public ValueTask<Transaction> CreateTransactionAsync(TransactionInfo transaction, CancellationToken token);

    public ValueTask<List<Transaction>> GetUserTransactionsAsync(long userId, CancellationToken token);
    public ValueTask<List<Transaction>> GetTransactionsAsync(CancellationToken token);

    public ValueTask<Transaction> ChangeTransactionStatusAsync(long transactionId, string sts,
        CancellationToken token);
    public ValueTask<bool> FinalizeTransactionsAsync(CancellationToken token);

}
