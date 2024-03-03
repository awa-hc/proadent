<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use app\Models\User;

class Appointment extends Model
{
    protected $fillable = [
        'user_id',
        'nro',
        'codigo',
        'type',
        'date',
        'reason',
        'user_auth_id',
        'status',
    ];

    public function user()
    {
        return $this->belongsTo(User::class);
    }

    const PENDIENTE = 1;
    const REALIZADA = 2;
    const ANULADA = 3;

    public function scopeByEstado($query, $estado)
    {
        if ($estado)  return $query->where('estado', $estado);
    }

    public function scopeByCodigo($query, $codigo)
    {
        if ($codigo)
        {
            return $query->where('codigo', 'LIKE', '%' . $codigo . '%');
        }
    }

    public function scopeByReason($query, $reason)
    {
        if ($reason)
        {
            return $query->where('reason', 'LIKE', '%' . $reason . '%');
        }
    }

    public function scopeByDate($query, $from, $to)
    {
        if ($from && $to)
        {

            return $query->whereBetween('date', [$from, $to]);
        }
    }

    public function scopeByUser($query, $user_id)
    {
        if ($user_id)
        {
            return $query->where('user_id', $user_id);
        }
    }

    public function getEstadoDescripcionAttribute()
    {
        $status_list = ['N/A', 'PENDIENTE', 'REALIZADA', 'ANULADA'];
        if (isset($status_list[$this->estado]))
        {
            return $status_list[$this->estado];
        }
        else
        {
            return '';
        }
    }

    public function getEstadoDescripcionAbreviadoAttribute()
    {
        $status_list = ['N/A', 'P', 'R', 'A'];
        if (isset($status_list[$this->estado]))
        {
            return $status_list[$this->estado];
        }
        else
        {
            return '';
        }
    }
}
