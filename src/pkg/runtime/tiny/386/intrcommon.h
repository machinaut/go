// this file gets #included over and over to workaround
// a problem in 8l when it processes JMPs between routines

  PUSHL DS
  PUSHL ES
  PUSHL FS
  PUSHL GS
  PUSHAL
  //MOVL  $(KDSEL),AX   - would have expanded to next line
  MOVL $((1<<3) | (0<<3) | 0), AX
  MOVW  AX,DS
  MOVW  AX,ES
  LEAL  0(SP),AX
  PUSHL AX
  CALL  runtime·trap(SB)
  POPL  AX
  POPAL
  POPL  GS
  POPL  FS
  POPL  ES
  POPL  DS
  ADDL  $8,SP /* error code and trap type */
  IRETL
