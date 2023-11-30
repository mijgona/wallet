using BusinessLogic;
using DataAccess;
using Microsoft.AspNetCore.Authorization;
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
    
    
    [Authorize]
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
    
    [Authorize]
    [HttpGet , Route("{userId}")]
    public async Task<IActionResult> GetUsersTransactions([FromRoute] long userId)
    {
        List<Transaction> res;
        try
        {
            res = await _transactionService.GetUserTransactionsAsync(userId, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    }    
    
    [Authorize]
    [HttpGet]
    public async Task<IActionResult> GetTransactions()
    {
        List<Transaction> res;
        try
        {
            res = await _transactionService.GetTransactionsAsync(new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    }
    
    [Authorize]
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