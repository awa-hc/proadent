<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\AuthController;
use App\Http\Controllers\RoleController;
use App\Http\Controllers\UserController;
use App\Http\Controllers\ProcedureController;
use App\Http\Controllers\PriceController;
use App\Http\Controllers\AppointmentController;
use App\Http\Controllers\AccountReceivableController;
use App\Http\Controllers\AccountReceivableDetailController;
use App\Http\Controllers\ReportController;

Route::middleware('auth:sanctum')->group(function ()
{

    Route::get('/user', function (Request $request)
    {
        return $request->user();
    });

    Route::post('logout', [AuthController::class, 'logout']);

    Route::post('/roles/store', [RoleController::class, 'store']);
    Route::get('/roles/index', [RoleController::class, 'index']);
    Route::get('/roles/{role}', [RoleController::class, 'show']);
    Route::put('/roles/{role}', [RoleController::class, 'update']);
    Route::delete('/roles/{role}', [RoleController::class, 'destroy']);

    Route::put('/users/{user}', [UserController::class, 'update']);
    Route::delete('/users/{user}', [UserController::class, 'destroy']);

    Route::post('/procedures/store', [ProcedureController::class, 'store']);
    Route::get('/procedures/index', [ProcedureController::class, 'index']);
    Route::get('/procedures/{procedure}', [ProcedureController::class, 'show']);
    Route::put('/procedures/{procedure}', [ProcedureController::class, 'update']);
    Route::delete('/procedures/{procedure}', [ProcedureController::class, 'destroy']);

    Route::post('/prices/store', [PriceController::class, 'store']);
    Route::get('/prices/index', [PriceController::class, 'index']);
    Route::get('/prices/{price}', [PriceController::class, 'show']);
    Route::put('/prices/{price}', [PriceController::class, 'update']);
    Route::delete('/prices/{price}', [PriceController::class, 'destroy']);

    Route::post('/appointments/store', [AppointmentController::class, 'store']);
    Route::get('/appointments/index', [AppointmentController::class, 'index']);
    Route::get('/appointments/{appointment}', [AppointmentController::class, 'show']);
    Route::put('/appointments/{appointment}', [AppointmentController::class, 'update']);
    Route::delete('/appointments/{appointment}', [AppointmentController::class, 'destroy']);

    Route::post('/account-receivables/store', [AccountReceivableController::class, 'store']);
    Route::get('/account-receivables/index', [AccountReceivableController::class, 'index']);
    Route::get('/account-receivables/{accountReceivable}', [AccountReceivableController::class, 'show']);
    Route::put('/account-receivables/{accountReceivable}', [AccountReceivableController::class, 'update']);
    Route::delete('/account-receivables/{accountReceivable}', [AccountReceivableController::class, 'destroy']);

    Route::post('/account-receivable-details/store', [AccountReceivableDetailController::class, 'store']);
    Route::get('/account-receivable-details/index', [AccountReceivableDetailController::class, 'index']);
    Route::get('/account-receivable-details/{accountReceivableDetail}', [AccountReceivableDetailController::class, 'show']);
    Route::put('/account-receivable-details/{accountReceivableDetail}', [AccountReceivableDetailController::class, 'update']);
    Route::delete('/account-receivable-details/{accountReceivableDetail}', [AccountReceivableDetailController::class, 'destroy']);

    Route::post('/reports/store', [ReportController::class, 'store']);
    Route::get('/reports/index', [ReportController::class, 'index']);
    Route::get('/reports/{report}', [ReportController::class, 'show']);
    Route::put('/reports/{report}', [ReportController::class, 'update']);
    Route::delete('/reports/{report}', [ReportController::class, 'destroy']);
});

Route::post('/login', [AuthController::class, 'login'])->withoutMiddleware('auth');
Route::post('/users/register', [UserController::class, 'register'])->withoutMiddleware('auth');
Route::post('/users/verify-token', [UserController::class, 'verifyToken'])->withoutMiddleware('auth');
Route::post('/users/send-email', [UserController::class, 'sendEmail'])->withoutMiddleware('auth');

Route::middleware('auth:sanctum')->get('/user', function (Request $request)
{
    return $request->user();
});
