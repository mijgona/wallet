﻿using BusinessLogic;
using DataAccess;
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
        Wallet res;
        try
        {
            res = await _walletService.CreateWalletAsync(request, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e);
        }
        
        return  Ok(res);
    } 
}