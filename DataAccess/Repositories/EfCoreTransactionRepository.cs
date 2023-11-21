﻿using System.Diagnostics;
using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public sealed class EfCoreTransactionRepository : ITransactionRepository
{
    private readonly WalletDbContext _db;

    public EfCoreTransactionRepository(WalletDbContext walletDbContext)
    {
        _db = walletDbContext;
    }
    
    public async ValueTask<Transaction> CreateAsync(Transaction transaction, CancellationToken token = default)
    {
        var res =  _db.Transactions.AddAsync(transaction, token).Result;
        await _db.SaveChangesAsync(token);
        
        return _db.Transactions.Find(res.Entity.Id) ?? new Transaction();
    }

    public async ValueTask<Transaction> UpdateStatusAsync(Transaction transaction, TransactionStatus newStatus,
        CancellationToken token = default)
    {
        transaction.Status = newStatus;
        var res =  _db.Transactions.Update(transaction);
        await _db.SaveChangesAsync(token);
        
        return _db.Transactions.Find(res.Entity.Id) ?? new Transaction();
    }

    public async ValueTask<string> GetStatusAsync(long transactionId, CancellationToken token = default)
    {
        var findAsync = await _db.Transactions.FindAsync(transactionId);
        Debug.Assert(findAsync != null, nameof(findAsync) + " != null");

        return _db.Transactions
            .Where(t => t.Id == transactionId)
            .Select(t => t.Status)
            .FirstOrDefault().ToString() ;
    }

    public async ValueTask<List<Transaction>> GetTransactionsAsync(long userId, CancellationToken token = default)
    {
        var userTransactions = _db.Transactions
            .Where(t => t.SourceWallet != null && t.SourceWallet.UserId == userId)
            .ToList();
        return await ValueTask.FromResult(userTransactions);
    }

    public ValueTask<ulong> GetTransactionsCountAsync(string userId, CancellationToken token = default)
    {
        throw new NotImplementedException();
    }

    public  async ValueTask<Transaction> GetTransactionByIdAsync(long transactionId, CancellationToken token = default)
    {
        return await ValueTask.FromResult(_db.Transactions
            .FirstOrDefault(t => t.Id == transactionId) ?? throw new InvalidOperationException());
    }
}