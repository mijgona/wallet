using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public interface ITransactionRepository
{
    ValueTask<Transaction> CreateAsync(Transaction transaction, CancellationToken token = default);
    ValueTask<bool> UpdateStatusAsync(Transaction transaction, string newStatus, CancellationToken token = default);
    ValueTask<string> GetStatusAsync(long transactionId, CancellationToken token = default);
    ValueTask<List<Transaction>> GetTransactionsAsync(long userId, CancellationToken token = default);
    ValueTask<ulong> GetTransactionsCountAsync(string userId, CancellationToken token = default);
}

