using BusinessLogic;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("[controller]")]
public class UserController : ControllerBase
{
    private readonly IUserService _userService;
    
    public UserController(IUserService userService)
    {
        _userService = userService;
    }
    
    [HttpPost]
    public async Task<IActionResult> Create([FromBody] UserInfo request)
    {
        var res =_userService.CreateUser(request).Result;
        return Ok(res);
    } 
}