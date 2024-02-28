<?php

namespace App\Http\Middleware;

use Closure;
use Illuminate\Auth\Middleware\Authenticate as Middleware;
use Illuminate\Http\Request;

class Authenticate extends Middleware
{
    /**
     * Handle an incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  \Closure  $next
     * @param  string|null  $guard
     * @return mixed
     */
    public function handle($request, Closure $next, ...$guards)
    {
        if ($this->authenticate($request, $guards) === false)
        {
            return response()->json(['message' => 'Unauthenticated'], 401);
        }

        return $next($request);
    }

    /**
     * Determine if the user is logged in to any of the given guards.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  array  $guards
     * @return bool
     */
    protected function authenticate($request, array $guards)
    {
        if (empty($guards))
        {
            $guards = [null];
        }

        foreach ($guards as $guard)
        {
            if ($this->auth->guard($guard)->check())
            {
                return true;
            }
        }

        return false;
    }
}
