using BusinessLogic;
using DataAccess;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace API.Controllers;

[ApiController]
[Route("[controller]")]
public class WorkerController : ControllerBase
{
    private readonly IWorkerService _workerService;
    private readonly ILogger<WorkerController> _logger;

    
    public WorkerController(IWorkerService workerService, ILogger<WorkerController> logger)
    {
        _workerService = workerService;
        _logger = logger;
    }
    
    [Authorize]
    [HttpPost]
    public Task<IActionResult> Create()
    {
        _logger.LogError("GET request received at {time}", DateTime.Now);

        try
        {
            _workerService.StartWorker();
        }
        catch (Exception e)
        {
            return Task.FromResult<IActionResult>(Conflict(e));
        }
        
        return  Task.FromResult<IActionResult>(Ok());
    } 
}