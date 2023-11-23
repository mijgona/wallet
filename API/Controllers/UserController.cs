using BusinessLogic;
using DataAccess;
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
    public async ValueTask<IActionResult> Create([FromBody] UserInfo request)
    {
        User res;
        try
        {
            res =await _userService.CreateUserAsync(request, new CancellationToken());

        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    } 
    
    [HttpGet]
    public async ValueTask<IActionResult> GetUserByUsername([FromQuery] string username)
    {
        User res;

        try
        {
            res = await _userService.GetUserByUsernameAsync(username, new CancellationToken());
        }
        catch (Exception e)
        {
            return Conflict(e) ;
        }
        return Ok(res);
    } 
}