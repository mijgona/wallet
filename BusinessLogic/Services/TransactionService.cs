using DataAccess;
using Microsoft.Extensions.Logging;

namespace BusinessLogic;

public sealed class TransactionService : ITransactionService
{
    private readonly ITransactionRepository _transactionRepository;
    
    private readonly ILogger<ITransactionService> _logger;

    public TransactionService(ITransactionRepository transactionRepository, ILogger<ITransactionService> logger)
    {
        _transactionRepository = transactionRepository;
        _logger = logger;
    }

    public ValueTask<Transaction> CreateTransactionAsync(TransactionInfo transaction, CancellationToken token)
    {

        Transaction newTransaction = new()
        {
            SourceWalletId = transaction.SourceWalletId,
            TargetWalletId = transaction.TargetWalletId,
            Status = TransactionStatus.Pending,
            Type = transaction.Type.ToTransactionTypeEnum(),
            Amount = transaction.Amount,
        };

        var result = _transactionRepository.CreateAsync(newTransaction, token);

        return result;
    }


    public async ValueTask<List<Transaction>> GetUserTransactionsAsync(long userId, CancellationToken token)
    {
        return await _transactionRepository.GetUserTransactionsAsync(userId, token);
        
    }

    public async ValueTask<List<Transaction>> GetTransactionsAsync(CancellationToken token)
    {
        return await _transactionRepository.GetTransactionsAsync(token);

    }

    public async ValueTask<Transaction> ChangeTransactionStatusAsync(long transactionId, string sts,
        CancellationToken token)
    {
        var status = sts.ToTransactionStatusEnum();
        var transaction = _transactionRepository.GetTransactionByIdAsync(transactionId, token).Result;

        if (transaction.Status == status)
        {
            return transaction;
        }

        return await _transactionRepository.UpdateStatusAsync(transaction, status, token);

    }

    public async ValueTask<bool> FinalizeTransactionsAsync(CancellationToken token)
    {
        bool changed = false;
        var transactions = await _transactionRepository.GetTransactionsByStatusAsync(TransactionStatus.Pending, token);

     
        _logger.LogError(transactions.Count.ToString() );
        
        if (transactions.Count > 0)
        {
            foreach (var transaction in transactions.Where(transaction => transaction.SourceWallet != null && (transaction.Amount < 0 || transaction.SourceWallet.Balance<0)))
            {
                changed = true;
                await _transactionRepository.UpdateStatusAsync(transaction, TransactionStatus.Cancelled, token);
            } 
            
            foreach (var transaction in transactions.Where(transaction => transaction.Amount > 0 && transaction is { SourceWalletId: > 0, TargetWalletId: > 0, SourceWallet.Balance: > 0 }))
            {
                changed = true;
                await _transactionRepository.UpdateStatusAsync(transaction, TransactionStatus.Success, token);
            }
        }

        return changed;

    }

}
