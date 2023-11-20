namespace DataAccess.Repositories;

public class TransactionRepository : ITransactionRepository
{
    public async ValueTask<bool> CreateAsync(Transaction transaction, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public ValueTask<bool> UpdateStatusAsync(Transaction transaction, string newStatus, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public ValueTask<string> GetStatusAsync(string transactionId, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public ValueTask<List<Transaction>> GetTransactionsAsync(string userId, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public ValueTask<ulong> GetTransactionsCountAsync(string userId, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }
}