using BusinessLogic;
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
        var res = await _transactionService.CreateTransactionAsync(request, new CancellationToken());
        return Ok(res);
    }  
    
    [HttpGet]
    public async Task<IActionResult> GetTransactions([FromHeader] long userId)
    {
        var res = await _transactionService.GetTransactionsAsync(userId, new CancellationToken());
        return Ok(res);
    }
}