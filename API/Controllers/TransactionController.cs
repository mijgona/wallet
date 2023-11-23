using BusinessLogic;
using DataAccess;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("[controller]")]
public class TransactionController : ControllerBase
{
    private readonly ITransactionService _transactionService;
    
    public TransactionController(ITransactionService transactionService)
    {
        _transactionService = transactionService;
    }
    
    
    [HttpPost]
    public async Task<IActionResult> CreateTransaction([FromBody] TransactionInfo request)
    {
        Transaction res;
        try
        {
            res = await _transactionService.CreateTransactionAsync(request, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    }  
    
    [HttpGet]
    public async Task<IActionResult> GetTransactions([FromHeader] long userId)
    {
        List<Transaction> res;
        try
        {
            res = await _transactionService.GetTransactionsAsync(userId, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    }
    
    [HttpPut]
    public async Task<IActionResult> ChangeTransactionStatus([FromHeader] string status, long transactionId)
    {
        Transaction res;
        try
        {
            res = await _transactionService.ChangeTransactionStatusAsync(transactionId, status, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    }
}