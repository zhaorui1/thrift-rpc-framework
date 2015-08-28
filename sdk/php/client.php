<?php

namespace addrservice\driver;

error_reporting(E_ALL);
require_once __DIR__.'/lib/Thrift/ClassLoader/ThriftClassLoader.php';
use Thrift\ClassLoader\ThriftClassLoader;
$GEN_DIR = realpath(dirname(__FILE__)).'/../../gen/gen-php';
$loader = new ThriftClassLoader();
$loader->registerNamespace('Thrift', __DIR__ . '/lib');
$loader->registerDefinition('addrservice', $GEN_DIR);
$loader->register();

use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\THttpClient;
use Thrift\Transport\TBufferedTransport;
use Thrift\Exception\TException;

try {

  $socket = new TSocket('localhost', 9090);
  $socket->setRecvTimeout(10000);
  $transport = new TBufferedTransport($socket, 1024, 1024);
  $protocol = new TBinaryProtocol($transport);
  $client = new \addrservice\AddressServiceClient($protocol);
  $transport->open();

  $req = new \addrservice\Request();
  $resp = new \addrservice\Response();

  $resp = $client->getAllAddress($req);

  echo $resp->code."\n";
  echo $resp->desc."\n";
  // echo $resp->data."\n";
  echo "Call finished.\n";

  $transport->close();

} catch (TException $tx) {
  print 'TException: '.$tx->getMessage()."\n";
}

?>