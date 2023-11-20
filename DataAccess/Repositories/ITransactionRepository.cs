using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public interface ITransactionRepository
{
    ValueTask<bool> CreateAsync(Transaction transaction, CancellationToken token = default);
    ValueTask<bool> UpdateStatusAsync(Transaction transaction, string newStatus, CancellationToken token = default);
    ValueTask<string> GetStatusAsync(string transactionId, CancellationToken token = default);
    ValueTask<List<Transaction>> GetTransactionsAsync(string userId, CancellationToken token = default);
    ValueTask<ulong> GetTransactionsCountAsync(string userId, CancellationToken token = default);
}

