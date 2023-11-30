using System.Diagnostics;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Internal;
using Microsoft.Extensions.Logging;

namespace DataAccess;

public sealed class EfCoreTransactionRepository : ITransactionRepository
{
    private readonly WalletDbContext _db;
    
    private readonly ILogger<ITransactionRepository> _logger;

    public EfCoreTransactionRepository(WalletDbContext walletDbContext, ILogger<ITransactionRepository> logger)
    {
        _db = walletDbContext;
        _logger = logger;
    }

    public async ValueTask<Transaction> CreateAsync(Transaction transaction, CancellationToken token = default)
    {
        var res = _db.Transactions.AddAsync(transaction, token).Result;
        await _db.SaveChangesAsync(token);

        return _db.Transactions.Find(res.Entity.Id) ?? new Transaction();
    }

    public async ValueTask<Transaction> UpdateStatusAsync(Transaction transaction, TransactionStatus newStatus,
        CancellationToken token = default)
    {
        _logger.LogError("Get on REPOSITORY");
        EntityEntry<Transaction> res;
        if (newStatus == TransactionStatus.Cancelled)
        {
            transaction.Status = newStatus;
        }
        else
        {
            transaction.SourceWallet.Balance -= transaction.Amount;
            transaction.TargetWallet.Balance += transaction.Amount;
            transaction.Status = newStatus;
        }

        res = _db.Transactions.Update(transaction);
        _db.Wallet.Update(transaction.SourceWallet);
        _db.Wallet.Update(transaction.TargetWallet);
        

        await _db.SaveChangesAsync(token);
        return await _db.Transactions.FindAsync( res.Entity.Id, token) ?? new Transaction();
    }

    public async ValueTask<string> GetStatusAsync(long transactionId, CancellationToken token = default)
    {
        var findAsync = await _db.Transactions.FindAsync(transactionId, token);
        Debug.Assert(findAsync != null, nameof(findAsync) + " != null");

        return findAsync.Status.ToString();
    }

    public async ValueTask<List<Transaction>> GetUserTransactionsAsync(long userId, CancellationToken token = default)
    {
        var userTransactions = _db.Transactions
            .Where(t => t.SourceWallet != null && t.SourceWallet.UserId == userId)
            .ToList();
        return await ValueTask.FromResult(userTransactions);
    }

    public async ValueTask<List<Transaction>>  GetTransactionsAsync(CancellationToken token = default)
    {
        var userTransactions = _db.Transactions.Where(t => t.Id != 0)
            .Include(t=>t.TargetWallet)
            .Where(t=>t.TargetWallet.Id == t.TargetWalletId) 
            .Include(t=>t.TargetWallet.User)
            .Where(t=>t.TargetWallet.User.Id == t.TargetWallet.UserId)
            .Include(t=>t.SourceWallet)
            .Where(t=>t.SourceWallet.Id == t.SourceWalletId) 
            .Include(t=>t.SourceWallet.User)
            .Where(t=>t.SourceWallet.User.Id == t.SourceWallet.UserId)
            .ToList();
        
        return await ValueTask.FromResult(userTransactions);
    }

    public async ValueTask<List<Transaction>> GetTransactionsByStatusAsync(TransactionStatus s,
        CancellationToken token = default)
    {
        var transactions = _db.Transactions
            .Where(t => t.Status == s)
            .Include(w => w.SourceWallet)
            .Where(w => w.SourceWallet.Id == w.SourceWalletId)
            .Include(t =>t.TargetWallet)
            .Where(t => t.TargetWallet.Id == t.TargetWalletId)
            .ToList();
        return await ValueTask.FromResult(transactions);
    }

    public async ValueTask<ulong> GetTransactionsCountAsync(long userId, CancellationToken token = default)
    {
        var res = _db.Transactions
            .Count(t => t.SourceWallet != null && t.SourceWallet.UserId == userId);
        return (ulong)await ValueTask.FromResult(res);
    }

    public async ValueTask<Transaction> GetTransactionByIdAsync(long transactionId, CancellationToken token = default)
    {
        return await ValueTask.FromResult(_db.Transactions
            .FirstOrDefault(t => t.Id == transactionId) ?? throw new InvalidOperationException());
    }
}