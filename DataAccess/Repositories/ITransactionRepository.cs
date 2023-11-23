using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public interface ITransactionRepository
{
    public ValueTask<Transaction> CreateAsync(Transaction transaction, CancellationToken token = default);

    public ValueTask<Transaction> UpdateStatusAsync(Transaction transaction, TransactionStatus newStatus,
        CancellationToken token = default);

    public ValueTask<string> GetStatusAsync(long transactionId, CancellationToken token = default);
    public ValueTask<List<Transaction>> GetTransactionsAsync(long userId, CancellationToken token = default);
    public ValueTask<ulong> GetTransactionsCountAsync(long userId, CancellationToken token = default);
    public ValueTask<Transaction> GetTransactionByIdAsync(long transactionId, CancellationToken token = default);

}

