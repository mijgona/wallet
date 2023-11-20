using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public sealed class EfCoreTransactionRepository : ITransactionRepository
{
    private readonly WalletDbContext _db;

    public EfCoreTransactionRepository(WalletDbContext walletDbContext)
    {
        _db = walletDbContext;
    }
    
    public async ValueTask<bool> CreateAsync(Transaction transaction, CancellationToken token = default)
    {
        await _db.Transactions.AddAsync(transaction, token);
        return await _db.SaveChangesAsync(token) > 0;
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