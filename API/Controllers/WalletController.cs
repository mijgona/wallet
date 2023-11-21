using BusinessLogic;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("[controller]")]
public class WalletController : ControllerBase
{
    private readonly IWalletService _walletService;
    
    public WalletController(IWalletService walletService)
    {
        _walletService = walletService;
    }
    
    [HttpPost]
    public async Task<IActionResult> Create([FromBody] WalletInfo request)
    {
        
        return Ok(_walletService.CreateWallet(request).Result);
    } 
}